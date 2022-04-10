import { serverPath } from './basePath';
import { definitions } from './generated-schema';

export type Post = definitions['Post'];

export async function GetAllPosts(basePath = serverPath): Promise<Post[]> {
  const raw = await fetch(`${basePath}/api/v1/posts/`);
  const posts: Post[] = await raw.json();
  return posts;
}

export async function GetOnePost(id: number, basePath = serverPath): Promise<Post> {
  const raw = await fetch(`${basePath}/api/v1/posts/${id}`);
  const post: Post = await raw.json();
  return post;
}
