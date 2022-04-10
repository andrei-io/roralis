import { ScreenParamsList } from '@/router/router';

export type RScreen = {
  name: string;
  routeName: keyof ScreenParamsList;
  iconName: string;
};

export const RDefaultScreens: RScreen[] = [
  {
    name: 'Posts',
    routeName: 'AllPosts',
    iconName: 'md-home-outline',
  },
  {
    name: 'Map',
    routeName: 'Dev',
    iconName: 'md-map-outline',
  },
  {
    name: 'Articles',
    routeName: 'AllArticles',
    iconName: 'md-pencil-outline',
  },
  {
    name: 'User',
    routeName: 'Profile',
    iconName: 'md-person-outline',
  },
];
