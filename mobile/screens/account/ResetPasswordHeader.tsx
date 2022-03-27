import { RText } from '@/components/ui/Text';
import colors from '@/shared/colors';
import { Ionicons } from '@expo/vector-icons';
import Constants from 'expo-constants';
import I18n from 'i18n-js';
import { FC } from 'react';
import { StyleSheet, View } from 'react-native';

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
export const ResetPasswordCodeHeader: FC = () => {
  return (
    <View style={headerStyles.headerContainer}>
      <View style={headerStyles.titleContainer}>
        <Ionicons name="close" size={25} color={colors.dark.accent} style={headerStyles.xIcon} />
        <RText text={I18n.t('resetPassword')} accent={true} variant="semiBold" size="large" />
      </View>
    </View>
  );
};
