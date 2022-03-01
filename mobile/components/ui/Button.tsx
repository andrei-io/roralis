import colors from '@shared/colors';
import React from 'react';
import { Pressable, StyleSheet, ViewStyle } from 'react-native';
import { RText } from './Text';

export interface IButtonProps {
  style?: ViewStyle;
  text?: string;
}

export const RButton: React.FC<IButtonProps> = ({ style = {}, text, children }) => {
  const combinedStyle = StyleSheet.compose(
    {
      backgroundColor: colors.dark.accent,
      justifyContent: 'center',
      alignItems: 'center',
      paddingHorizontal: 50,
      paddingVertical: 16,
      borderRadius: 200,
    } as ViewStyle,
    style,
  );
  return (
    <Pressable style={combinedStyle}>
      {text && <RText text={text} />}
      {children}
    </Pressable>
  );
};
