import colors from '@/shared/colors';
import I18n from 'i18n-js';
import { FC } from 'react';
import { Image, StyleSheet, View } from 'react-native';
import { RText } from '../ui/Text';

export interface IPostProps {
  title: string;
  description: string;
  time?: string;
}

const styles = StyleSheet.create({
  bigContainer: {
    width: '100%',
    flexDirection: 'row',
    alignItems: 'flex-start',
  },
  tinyImage: {
    width: 50,
    height: 50,
    borderRadius: 8,
    marginTop: 5,
    marginHorizontal: 10,
  },
  textContainer: {
    flexDirection: 'column',
    flex: 1,
    paddingRight: 10,
  },
  headerContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  separator: {
    backgroundColor: colors.dark.lightGray,
    marginTop: 10,
    height: 1,
  },
});

export const RPostNormal: FC<IPostProps> = ({ title, description, time = '?' }) => {
  return (
    <View style={styles.bigContainer}>
      <Image style={styles.tinyImage} source={require('@/assets/post-placeholder.png')} />
      <View style={styles.textContainer}>
        <View style={styles.headerContainer}>
          <RText text={title} style={{ color: colors.dark.mediumGray }} variant="semiBold" />
          <RText
            text={time + ' ' + I18n.t('ago')}
            style={{ color: colors.dark.mediumGray }}
            variant="semiBold"
          />
        </View>
        <RText text={description} style={{ color: colors.dark.mediumGray }} />
        <View style={styles.separator} />
      </View>
    </View>
  );
};
