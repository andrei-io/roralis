import { GenericResponse, serverPath } from './constants';
import { definitions, responses } from './generated-schema';

export type User = definitions['User'];

type SignInSucces = responses['SignInSucces']['schema'];
type SignUpSucces = responses['SignUpSucces']['schema'];

export async function GetOneUser(
  id: number,
  abort?: AbortController,
  basePath = serverPath,
): Promise<User> {
  const raw = await fetch(`${basePath}/api/v1/users/${id}`, { signal: abort?.signal });
  const post: User = await raw.json();
  return post;
}

export async function SignIn(
  email: string,
  password: string,
  abort?: AbortController,
  basePath = serverPath,
): Promise<string> {
  const requestBody = {
    Email: email,
    Password: password,
  };
  const raw = await fetch(`${basePath}/api/v1/users/signin`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    signal: abort?.signal,
    body: JSON.stringify(requestBody),
  });

  if (raw.status == 401) throw new Error('Your password or email are incorrect');
  if (raw.status == 404) throw new Error('Account not found');

  if (!raw.ok) {
    const error: GenericResponse = await raw.json();
    throw new Error(`Request failed: ${error.Message} `);
  }

  const json: SignInSucces = await raw.json();

  if (!json.Token) throw new Error('Invalid request');
  return json.Token;
}

export async function SignUp(
  name: string,
  email: string,
  password: string,
  abort?: AbortController,
  basePath = serverPath,
): Promise<SignUpSucces> {
  const requestBody = {
    Name: name,
    Email: email,
    Password: password,
  };
  const raw = await fetch(`${basePath}/api/v1/users/signup`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    signal: abort?.signal,
    body: JSON.stringify(requestBody),
  });

  if (raw.status == 409) throw new Error('Your email or username is already used');

  if (!raw.ok) {
    const error: GenericResponse = await raw.json();
    throw new Error(`Request failed: ${error.Message} `);
  }

  const json: SignUpSucces = await raw.json();

  if (!json.Token) throw new Error('Invalid request');
  return json;
}
