import colors from '@/shared/colors';
import I18n from 'i18n-js';
import { FC } from 'react';
import { Image, StyleSheet, View } from 'react-native';
import { RText } from '../ui/Text';

export interface IArticleProps {
  title: string;
  description: string;
  time?: string;
}

const styles = StyleSheet.create({
  container: {
    width: '90%',
  },
  image: {
    width: '100%',
    resizeMode: 'center',
    height: 240,
    borderRadius: 8,
  },
});

export const RArticle: FC<IArticleProps> = ({ title, description, time = '?' }) => {
  return (
    <View style={styles.container}>
      <Image source={require('@/assets/corn.jpg')} style={styles.image} />
      <RText text={title} style={{ color: colors.dark.mediumGray }} variant="semiBold" />
      <RText text={description} style={{ color: colors.dark.mediumGray }} variant="regular" />
      <RText
        text={time + ' ' + I18n.t('ago')}
        style={{ color: colors.dark.mediumGray, alignSelf: 'flex-end' }}
        variant="regular"
      />
    </View>
  );
};
