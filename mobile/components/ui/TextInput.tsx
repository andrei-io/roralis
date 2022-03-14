import colors from '@/shared/colors';
import React from 'react';
import { StyleSheet, TextInput, TextStyle, ViewStyle } from 'react-native';

interface ITextInputProps {
  style?: any;
  placeholder?: string;
}

export const RTextInput: React.FC<ITextInputProps> = ({ style = {}, placeholder }) => {
  const combinedStyle = StyleSheet.compose(
    {
      backgroundColor: colors.dark.lightGray,
      minWidth: 100,
      paddingHorizontal: 10,
      fontFamily: 'Inter_500Medium',
    } as ViewStyle & TextStyle,
    style,
  );
  return <TextInput style={combinedStyle} placeholder={placeholder} />;
};
