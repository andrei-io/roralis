import colors from '@/shared/colors';
import React from 'react';
import { Keyboard, StyleSheet, TextInput, ViewStyle } from 'react-native';

interface IVerificationCodeProps {
  onChange(code: string): void;
  codeLength?: number;
  style?: ViewStyle;
}

export const RVerificationCode: React.FC<IVerificationCodeProps> = ({
  codeLength = 4,
  onChange,
  style,
}) => {
  const combinedStyle = StyleSheet.compose(
    {
      minWidth: codeLength * 40 + (codeLength - 1) * 10,
      borderWidth: 2,
      borderColor: colors.dark.accent,
      color: colors.dark.accent,
      textAlign: 'center',
      backgroundColor: 'transparent',
      fontSize: 50,
      letterSpacing: 15,
    } as ViewStyle,
    style,
  );

  return (
    <TextInput
      style={combinedStyle}
      onChangeText={(value) => {
        if (value.length == codeLength) Keyboard.dismiss();
        onChange(value);
      }}
      selectionColor={colors.dark.accent}
      keyboardType="number-pad"
      placeholder={'0'.repeat(codeLength)}
      // Adding 22 at the end of a rgb hex string specifies opacity
      placeholderTextColor={colors.dark.mediumGray + '22'}
      maxLength={codeLength}
    />
  );
};
