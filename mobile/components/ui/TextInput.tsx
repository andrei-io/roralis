import colors from '@/shared/colors';
import React from 'react';
import { StyleSheet, TextInput, TextStyle, ViewStyle } from 'react-native';
import { TextVariant } from './Text';

interface ITextInputProps {
  style?: any;
  placeholder?: string;
}

export const RTextInput: React.FC<ITextInputProps> = ({ style, placeholder }) => {
  const combinedStyle = StyleSheet.compose(
    {
      backgroundColor: colors.dark.lightGray,
      paddingHorizontal: 10,
      paddingVertical: 5,
      borderRadius: 8,
      fontFamily: TextVariant.medium,
    } as ViewStyle & TextStyle,
    style,
  );
  return (
    <TextInput
      style={combinedStyle}
      placeholder={placeholder}
      selectionColor={colors.dark.accent}
    />
  );
};
