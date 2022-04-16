import { User } from '@/restapi/UserAPI';
import { getItemAsync } from 'expo-secure-store';
import keys from './keys';

export async function isLoggedIn(): Promise<boolean> {
  return ((await getItemAsync(keys.user)) ?? '') != '';
}

export async function getUserCache(): Promise<User> {
  const raw = await getItemAsync(keys.user);
  if (!raw) throw new Error('No user cache found');
  return JSON.parse(raw);
}
