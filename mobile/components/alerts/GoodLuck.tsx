import colors from '@/shared/colors';
import I18n from 'i18n-js';
import { FC } from 'react';
import { StyleSheet, View } from 'react-native';
import { RButton } from '../ui/Button';
import { RText } from '../ui/Text';

export interface IGoodLuckPopupProps {
  onFinish?(): void;
}

const styles = StyleSheet.create({
  container: {
    width: '90%',
    height: '60%',
    paddingTop: 25,
    paddingHorizontal: 25,
    borderRadius: 8,
    backgroundColor: colors.dark.white,
    alignItems: 'center',
    justifyContent: 'space-evenly',
  },
});

export const RGoodLuckPopup: FC<IGoodLuckPopupProps> = ({ onFinish }) => {
  return (
    <View style={styles.container}>
      <RText text={I18n.t('goodLuckTitle')} size="large" variant="semiBold" />
      <RText
        text={I18n.t('goodLuckMessage')}
        size="semiLarge"
        variant="medium"
        style={{
          color: colors.dark.darkGray,
        }}
      />
      <RButton
        text={I18n.t('continue')}
        style={{
          width: '100%',
        }}
        onClick={onFinish}
      />
    </View>
  );
};
