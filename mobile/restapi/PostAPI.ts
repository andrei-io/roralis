import { serverPath } from './basePath';
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
