import { serverPath } from './basePath';
import { definitions, responses } from './generated-schema';

export type User = definitions['User'];

type SignInSucces = responses['SignInSucces']['schema'];

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
    throw new Error('Request failed');
  }

  const json: SignInSucces = await raw.json();

  if (!json.Token) throw new Error('Invalid request');
  return json.Token;
}
