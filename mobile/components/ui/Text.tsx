import colors from '@/shared/colors';
import React from 'react';
import { StyleSheet, Text, TextStyle } from 'react-native';

export enum TextSize {
  large = 30,
  semiLarge = 24,
  normal = 16,
  small = 10,
}

export enum TextVariant {
  regular = 'Inter_400Regular',
  medium = 'Inter_500Medium',
  semiBold = 'Inter_600SemiBold',
  bold = 'Inter_700Bold',
}

export interface ITextProps {
  style?: TextStyle;
  size?: keyof typeof TextSize;
  accent?: boolean;
  variant?: keyof typeof TextVariant;
  text: string;
}

export const RText: React.FC<ITextProps> = ({
  size = 'normal',
  accent = false,
  text,
  variant = 'regular',
  style,
}) => {
  const combinedStyle = StyleSheet.compose(
    {
      color: accent ? colors.dark.accent : colors.dark.black,
      fontSize: TextSize[size],
      fontFamily: TextVariant[variant],
    } as TextStyle,
    style,
  );
  return <Text style={combinedStyle}>{text}</Text>;
};
