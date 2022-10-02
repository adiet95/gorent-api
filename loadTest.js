import http from 'k6/http';
import { check, group, sleep, fail } from 'k6';

export const options = {
  stages: [{ target: 3, duration: '10s' }],
  thresholds: {
    'http_req_duration': ['p(95)<500', 'p(99)<1500'],
    'http_req_duration{name:PublicCrocs}': ['avg<400'],
    'http_req_duration{name:Create}': ['avg<600', 'max<1000'],
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
  // register a new user and authenticate via a Bearer token.
//   const res = http.post(`${BASE_URL}/user/register/`, {
//     first_name: 'Crocodile',
//     last_name: 'Owner',
//     username: EMAIL,
//     password: PASSWORD,
//   });

//   check(res, { 'created user': (r) => r.status === 201 });

//Login and auth vvia bearer token
  const loginRes = http.post(`${BASE_URL}`, {
    email: EMAIL,
    password: PASSWORD,
  });

  const authToken = loginRes.json("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluIiwiUm9sZSI6ImFkbWluIiwiZXhwIjoxNjY0NzU2MjMxfQ.qq38p37jx35Ku2gG81_U2ayckJNIAfajlQJcAS8lCIU");
  check(authToken, { 'logged in successfully': () => authToken !== '' });

  return authToken;
}

export default (authToken) => {
  const requestConfigWithTag = (tag) => ({
    headers: {
      Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluIiwiUm9sZSI6ImFkbWluIiwiZXhwIjoxNjY0NzU2NDIzfQ.HkuMJItzCV37xgStCTHyzvOjMzqQsMoHMtHyeOeIcNw `,
    },
    tags: Object.assign(
      {},
      {
        name: 'Private',
      },
      tag
    ),
  });

//   group('Public endpoints', () => {
//     // call some public endpoints in a batch
//     const responses = http.batch([
//       ['GET', `${BASE_URL}/public/crocodiles/1/`, null, { tags: { name: 'PublicCrocs' } }],
//       ['GET', `${BASE_URL}/public/crocodiles/2/`, null, { tags: { name: 'PublicCrocs' } }],
//       ['GET', `${BASE_URL}/public/crocodiles/3/`, null, { tags: { name: 'PublicCrocs' } }],
//       ['GET', `${BASE_URL}/public/crocodiles/4/`, null, { tags: { name: 'PublicCrocs' } }],
//     ]);

//     const ages = Object.values(responses).map((res) => res.json('age'));

//     // Functional test: check that all the public crocodiles are older than 5
//     check(ages, {
//       'Crocs are older than 5 years of age': Math.min(...ages) > 5,
//     });
//   });

  group('Create and modify crocs', () => {
    let URL = `${BASE_URL}user/`;

    group('Create crocs', () => {
      const payload = {
        email: `${randomString(10)}`,
        password: 'user',
        role: 'user',
      };

      const res = http.post(URL, payload, requestConfigWithTag({ name: 'Create' }));

      if (check(res, { 'User created correctly': (r) => r.status === 200 })) {
        URL = `${URL}${res.json('id')}/`;
      } else {
        console.log(`Unable to create a user ${res.status} ${res.body}`);
        return;
      }
    });

    group('Update croc', () => {
      const payload = { 
        password: 'NewPass',
        address: 'Jakarta',
        phone: '+62'
     };

      const res = http.put(URL, payload,requestConfigWithTag({ name: 'Update' }));

      const isSuccessfulUpdate = check(res, {
        'Update worked': () => res.status === 200,
        'Updated is correct': () => res.json('password') === 'NewPass',
      });

      if (!isSuccessfulUpdate) {
        console.log(`Unable to update the croc ${res.status} ${res.body}`);
        return;
      }
    });

    const delRes = http.del(URL, null, requestConfigWithTag({ name: 'Delete' }));

    const isSuccessfulDelete = check(null, {
      'Croc was deleted correctly': () => delRes.status === 200,
    });

    if (!isSuccessfulDelete) {
      console.log(`Croc was not deleted properly`);
      return;
    }
  });

  sleep(1);
};