import colors from '@/shared/colors';
import { Ionicons } from '@expo/vector-icons';
import React, { FC } from 'react';
import { StyleSheet, View } from 'react-native';

export type RScreen = {
  name: string;
  routeName: string;
  iconName: string;
  checked?: boolean;
};

interface INavigationBarProps {
  screens?: RScreen[];
}

export const RDefaultScreens: RScreen[] = [
  {
    name: 'Posts',
    routeName: '',
    iconName: 'md-home-outline',
    checked: true,
  },
  {
    name: 'Map',
    routeName: '',
    iconName: 'md-map-outline',
  },
  {
    name: 'Articles',
    routeName: '',
    iconName: 'md-pencil-outline',
  },
  {
    name: 'User',
    routeName: '',
    iconName: 'md-person-outline',
  },
];

const styles = StyleSheet.create({
  container: {
    position: 'absolute',
    bottom: 0,
    height: '10%',
    width: '100%',
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center',
  },
  icon: {
    backgroundColor: colors.dark.white,
    justifyContent: 'center',
    alignItems: 'center',
    width: 40,
    height: 40,
    borderRadius: 40,
  },
  selectedIcon: {
    backgroundColor: colors.dark.accent,
    justifyContent: 'center',
    alignItems: 'center',
    width: 40,
    height: 40,
    borderRadius: 40,
  },
});

export const RNavigationBar: FC<INavigationBarProps> = ({ screens = RDefaultScreens }) => {
  // Ignoring typescript beacuse i can't find the correct type for Ionicons.name
  // I have tried:
  //     - checking the type definitions
  //     - using keyof typeof Ionicons.name

  return (
    <View style={styles.container}>
      {screens.map((screen, i) => (
        <View style={screen.checked ? styles.selectedIcon : styles.icon} key={i}>
          {/* @ts-ignore */}
          <Ionicons name={screen.iconName} size={24} color={colors.dark.black} />
        </View>
      ))}
    </View>
  );
};
