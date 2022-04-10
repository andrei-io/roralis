import colors from '@/shared/colors';
import { StringToColor } from '@/shared/generateColor';
import React from 'react';
import { StyleSheet, View, ViewStyle } from 'react-native';
import { RText } from '../ui/Text';

interface IUserPhotoProps {
  name: string;
}

export const iconSize = 200;

const styles = StyleSheet.create({
  container: {
    position: 'relative',
    top: (-3 / 4) * iconSize,
    width: iconSize,
    height: iconSize,
    borderRadius: 1000,
    justifyContent: 'center',
    alignItems: 'center',
    borderColor: colors.dark.white,
    borderWidth: 4,
  },
});

export const RUserPhoto: React.FC<IUserPhotoProps> = ({ name }) => {
  // This takes max 2 initials from the name and
  // Because there is always an extra space at the end we add one at the beginnig for centering
  let initials = name
    .split('')
    .filter((c) => c.toUpperCase() == c && c != ' ')
    .filter((_, i) => i < 2)
    .reduce((str, c) => str + c + '. ', ' ');
  if (initials == ' ') initials = ' A. A. ';

  return (
    <View
      style={{
        ...styles.container,
        ...({
          backgroundColor: StringToColor(name),
        } as ViewStyle),
      }}
    >
      <RText
        text={initials}
        size={'extraLarge'}
        style={{
          color: colors.dark.white,
        }}
        variant="semiBold"
      />
    </View>
  );
};
