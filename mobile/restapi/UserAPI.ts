import { serverPath } from './basePath';
import { definitions } from './generated-schema';

export type User = definitions['User'];

export async function GetOneUser(id: number, basePath = serverPath): Promise<User> {
  const raw = await fetch(`${basePath}/api/v1/users/${id}`);
  const post: User = await raw.json();
  return post;
}
