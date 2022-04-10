import { RText } from '@/components/ui/Text';
import colors from '@/shared/colors';
import { Ionicons } from '@expo/vector-icons';
import Constants from 'expo-constants';
import { FC } from 'react';
import { Pressable, StyleSheet, View } from 'react-native';

interface ISimpleHeaderProps {
  title: string;
  onX?(): void;
}

const headerStyles = StyleSheet.create({
  headerContainer: {
    backgroundColor: colors.dark.background,
    width: '100%',
    height: '13%',
  },
  xIcon: {
    paddingHorizontal: 10,
    position: 'absolute',
    left: '3%',
  },
  titleContainer: {
    top: Constants.statusBarHeight,
    flexDirection: 'row',
    justifyContent: 'center',
    alignItems: 'center',
  },
});
export const SimpleHeader: FC<ISimpleHeaderProps> = ({ title, onX }) => {
  return (
    <View style={headerStyles.headerContainer}>
      <View style={headerStyles.titleContainer}>
        <Pressable onPress={onX} style={headerStyles.xIcon} android_disableSound={true}>
          <Ionicons name="close" size={25} color={colors.dark.accent} />
        </Pressable>
        <RText text={title} accent={true} variant="semiBold" size="large" />
      </View>
    </View>
  );
};
