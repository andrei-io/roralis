import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { Ionicons } from '@expo/vector-icons';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import React, { FC } from 'react';
import { Pressable, StyleSheet, View } from 'react-native';
import { RDefaultScreens, RScreen } from '../../shared/screens';

interface INavigationBarProps {
  screens?: RScreen[];
  navigation: NativeStackNavigationProp<ScreenParamsList, keyof ScreenParamsList>;
  focused: keyof ScreenParamsList;
}

const styles = StyleSheet.create({
  container: {
    position: 'absolute',
    bottom: 0,
    height: '10%',
    width: '100%',
    flexDirection: 'row',
    justifyContent: 'space-around',
    alignItems: 'center',
    backgroundColor: colors.dark.background,
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

export const RNavigationBar: FC<INavigationBarProps> = ({
  screens = RDefaultScreens,
  focused,
  navigation,
}) => {
  // Ignoring typescript beacuse i can't find the correct type for Ionicons.name
  // I have tried:
  //     - checking the type definitions
  //     - using keyof typeof Ionicons.name

  return (
    <View style={styles.container}>
      {screens.map((screen, i) => (
        <Pressable
          style={screen.routeName == focused ? styles.selectedIcon : styles.icon}
          onPress={() => {
            navigation.push(screen.routeName);
          }}
          android_disableSound
          key={i}
        >
          {/* @ts-ignore */}
          <Ionicons name={screen.iconName} size={24} color={colors.dark.black} />
        </Pressable>
      ))}
    </View>
  );
};
