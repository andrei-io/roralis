import { GenericResponse, serverPath } from './constants';
import { definitions } from './generated-schema';

export type Post = definitions['Post'];

export async function GetAllPosts(abort?: AbortController, basePath = serverPath): Promise<Post[]> {
  const raw = await fetch(`${basePath}/api/v1/posts/`, { signal: abort?.signal });
  const posts: Post[] = await raw.json();
  return posts;
}

export async function GetOnePost(
  id: number,
  abort?: AbortController,
  basePath = serverPath,
): Promise<Post> {
  const raw = await fetch(`${basePath}/api/v1/posts/${id}`, { signal: abort?.signal });
  const post: Post = await raw.json();
  return post;
}

export async function CreatePost(
  post: Post,
  jwt: string,
  abort?: AbortController,
  basePath = serverPath,
) {
  var requestOptions = {
    method: 'POST',
    redirect: 'follow',
    headers: {
      'Content-Type': 'application/json',
      Authorization: jwt,
    },
    signal: abort?.signal,
    body: JSON.stringify(post),
  };

  const raw = await fetch(`${basePath}/api/v1/posts`, requestOptions);

  if (raw.status == 401) throw new Error('You are not logged in');
  if (raw.status == 403) throw new Error(`You don't have permission`);
  if (!raw.ok) {
    const error: GenericResponse = await raw.json();
    throw new Error(error.Message);
  }
}
