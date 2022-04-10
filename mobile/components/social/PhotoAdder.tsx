import colors from '@/shared/colors';
import { Ionicons } from '@expo/vector-icons';
import { FC } from 'react';
import { StyleSheet, View } from 'react-native';
import { RText } from '../ui/Text';

const styles = StyleSheet.create({
  container: {
    backgroundColor: colors.dark.white,
    width: '100%',
    paddingLeft: 10,
    paddingRight: 30,
    borderRadius: 8,
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  plus: {
    fontSize: 40,
    color: colors.dark.green,
  },
  minus: {
    fontSize: 40,
    color: colors.dark.red,
  },
});

export const RPhotoAdder: FC = () => {
  return (
    <View style={styles.container}>
      <Ionicons name="md-camera-outline" size={40} />
      <RText text="+" style={styles.plus} />
      <RText text="-" style={styles.minus} />
    </View>
  );
};
