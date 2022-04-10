import colors from '@/shared/colors';
import { StyleSheet, TextInput, TextStyle, ViewStyle } from 'react-native';
import { TextVariant } from './Text';

interface ITextInputProps {
  style?: any;
  placeholder?: string;
  password?: boolean;
  onChange?(text: string): void;
}

export const RTextInput: React.FC<ITextInputProps> = ({
  style,
  onChange,
  placeholder,
  password = false,
}) => {
  const combinedStyle = StyleSheet.compose(
    {
      backgroundColor: colors.dark.lightGray,
      paddingHorizontal: 20,
      paddingVertical: 10,
      borderRadius: 8,
      fontFamily: TextVariant.medium,
    } as ViewStyle & TextStyle,
    style,
  );
  return (
    <TextInput
      style={combinedStyle}
      placeholder={placeholder}
      onChangeText={onChange}
      selectionColor={colors.dark.accent}
      secureTextEntry={password}
    />
  );
};
