import { User } from '@/restapi/UserAPI';
import { getItemAsync, setItemAsync } from 'expo-secure-store';
import keys from './keys';

export async function isLoggedIn(): Promise<boolean> {
  return ((await getItemAsync(keys.user)) ?? '') != '';
}

export async function getUserCache(): Promise<User> {
  const raw = await getItemAsync(keys.user);
  if (!raw) throw new Error('No user cache found');
  return JSON.parse(raw);
}

export async function setUserCache(user: User): Promise<void> {
  await setItemAsync(keys.user, JSON.stringify(user));
}

export async function setToken(token: string): Promise<void> {
  await setItemAsync(keys.token, token);
}
