import colors from '@/shared/colors';
import React from 'react';
import { StyleSheet, Text, TextStyle } from 'react-native';

// type TextSize = 'large' | 'normal' | 'small';

enum TextSize {
  large = 30,
  normal = 16,
  small = 10,
}

export interface ITextProps {
  style?: TextStyle;
  size?: TextSize;
  accent?: boolean;
  text: string;
}

export const RText: React.FC<ITextProps> = ({
  size = 'normal',
  accent = false,
  text,
  style = {},
}) => {
  const combinedStyle = StyleSheet.compose(
    {
      color: accent ? colors.dark.accent : colors.dark.black,
      fontSize: TextSize[size],
      fontFamily: 'Inter_600SemiBold',
    } as TextStyle,
    style,
  );
  return <Text style={combinedStyle}>{text}</Text>;
};
