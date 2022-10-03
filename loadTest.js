import http from 'k6/http';
import { check, group, sleep, fail } from 'k6';

export const options = {
  ext: {
    loadimpact: {
      apm: [
        {
          provider: 'prometheus',
          remoteWriteURL: 'https://prometheus-prod-10-prod-us-central-0.grafana.net/api/prom/push',
          credentials: {
            username: '593457',
            password: 'eyJrIjoiNWZjODQ0ODQwMDcyMDY2YTQyODMxYTFmZTk3MTkxZTZlN2Y2ZjJkNyIsIm4iOiJhZGlldDIyIiwiaWQiOjcxOTA4OH0=',
          },
          // optional parameters
          metrics: [
            'vus',
            'http_req_duration',
            'my_rate_metric',
            'my_gauge_metric',
            'data_sent',
            'data_received',
            // create a metric by counting HTTP responses with status 500
            {
              sourceMetric: 'http_reqs{status="500"}',
              targetMetric: 'k6_http_server_errors_count',
              keepTags: ['name', 'method', 'status'],
            },
            {
              sourceMetric: 'http_reqs{status="500"}',
              targetMetric: 'k6_http_server_errors_count',
              keepTags: ['scenario', 'group', 'name', 'method'],
            },
          ],
          // for advanced metric configuration see example belod
          includeDefaultMetrics: true,
          includeTestRunId: false,
        },
      ],

      projectID: 3603436,
      // Test runs with the same name groups test runs together
      name: "Gorent App"
    }
  },
  stages: [
            { target: 5, duration: '1s' },
            { target: 50, duration: '10s' },
            { target: 0, duration: '1s' }
        ],
  thresholds: {
    'http_req_duration': ['p(95)<500', 'p(99)<1500'],
    'http_req_duration{name:PublicCrocs}': ['avg<400'],
    'http_req_duration{name:Create}': ['avg<600', 'max<1000'],
    'http_req_duration{name:Update}': ['avg<600', 'max<1000'],
    'http_req_duration{name:Get}': ['avg<600', 'max<1000'],
  },
};

function randomString(length, charset = '') {
  if (!charset) charset = 'abcdefghijklmnopqrstuvwxyz';
  let res = '';
  while (length--) res += charset[(Math.random() * charset.length) | 0];
  return res;
}

const EMAIL = 'admin';
const PASSWORD = 'admin';
const BASE_URL = 'https://gorent-api.herokuapp.com/';

export function setup() {
//   register a new user and authenticate via a Bearer token.
  const res = http.post(`${BASE_URL}register`, {
    username: EMAIL,
    password: PASSWORD,
  });

  check(res, { 'created user': (r) => r.status === 200 });

//Login and auth vvia bearer token
  const loginRes = http.post(`${BASE_URL}`, {
    email: EMAIL,
    password: PASSWORD,
  });

  const authToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluIiwiUm9sZSI6ImFkbWluIiwiZXhwIjoxNjY0ODE3NjE3fQ.Hu0Drur64V7Uq-pzgz6aI2z6z1yj646qQ6Eah4iyP40";

  check(authToken, { 'logged in successfully': () => authToken !== '' });

  return authToken;
}

export default (authToken) => {
  const requestConfigWithTag = (tag) => ({
    headers: {
      Authorization: `Bearer ${authToken} `,
    },
    tags: Object.assign(
      {},
      {
        name: 'Private',
      },
      tag
    ),
  });

  group('Create and Get User', () => {
    let URL = `${BASE_URL}user/`;

    let ranEmail = `${randomString(10)}`

    group('Get User', () => {  
        const res = http.get(URL, requestConfigWithTag({ name: 'Get' }));
  
        if (check(res, { 'Get User data correctly': (r) => r.status === 200 })) {
          URL = `${URL}`;
        } else {
          console.log(`Unable to get a user ${res.status}`);
          return;
        }
      });

    group('Create User', () => {
      const payload = {
        email: ranEmail,
        password: 'user',
        role: 'user',
      };

      const res = http.post(URL, payload, requestConfigWithTag({ name: 'Create' }));

      if (check(res, { 'User created correctly': (r) => r.status === 200 })) {
        URL = `${URL}`;
      } else {
        console.log(`Unable to create a user ${res.status} ${res.body}`);
        return;
      }
    });

    group('Update croc', () => {
      const payload = {address: 'Jakarta'};

      const res = http.put(URL, payload,requestConfigWithTag({ name: 'Update' }));

      const isSuccessfulUpdate = check(res, {
        'Update worked': () => res.status === 200
      });

      if (!isSuccessfulUpdate) {
        console.log(`Unable to update the croc ${res.status} ${res.body}`);
        return;
      }
    });
    const payload = {
        email: `?email=${ranEmail}`
      };

  });
  sleep(1);
};