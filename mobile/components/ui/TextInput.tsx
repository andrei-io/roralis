import colors from "@/shared/colors";
import { FC } from "react";
import { StyleSheet, TextInput, TextStyle, ViewStyle } from "react-native";
import { TextVariant } from "./Text";

interface ITextInputProps {
  style?: any;
  placeholder?: string;
  password?: boolean;
}

export const RTextInput: FC<ITextInputProps> = ({
  style,
  placeholder,
  password = false,
}) => {
  const combinedStyle = StyleSheet.compose(
    {
      backgroundColor: colors.dark.lightGray,
      paddingHorizontal: 10,
      paddingVertical: 5,
      borderRadius: 8,
      fontFamily: TextVariant.medium,
    } as ViewStyle & TextStyle,
    style
  );
  return (
    <TextInput
      style={combinedStyle}
      placeholder={placeholder}
      selectionColor={colors.dark.accent}
      secureTextEntry={password}
    />
  );
};
