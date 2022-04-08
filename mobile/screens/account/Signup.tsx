import { RButton } from "@/components/ui/Button";
import { RCheckbox } from "@/components/ui/Checkbox";
import { RText } from "@/components/ui/Text";
import { RTextInput } from "@/components/ui/TextInput";
import { ScreenParamsList } from "@/router/router";
import colors from "@/shared/colors";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { StatusBar } from "expo-status-bar";
import I18n from "i18n-js";
import React from "react";
import { StyleSheet, View } from "react-native";
import { SimpleHeader } from "../header/SimpleHeader";

interface ISignuprops {}
type RouterProps = NativeStackScreenProps<ScreenParamsList, "Login">;
const styles = StyleSheet.create({
  container: {
    backgroundColor: colors.dark.background,
    flex: 1,
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
  },
  button: {
    width: "90%",
    marginVertical: "20%",
  },
  mail: {
    width: "90%",
    marginBottom: "20%",
    height: "6%",
  },
  password: {
    width: "90%",
    marginBottom: "20%",
    height: "6%",
  },
  newsletter: {
    flexDirection: "row",
  },
});
const SignupScreen: React.FC<ISignuprops & RouterProps> = ({ navigation }) => {
  return (
    <>
      <SimpleHeader title={I18n.t("signup")} />
      <View style={styles.container}>
        <RTextInput style={styles.mail} placeholder={I18n.t("name")} />
        <RTextInput style={styles.mail} placeholder={I18n.t("email")} />
        <RTextInput
          style={styles.password}
          placeholder={I18n.t("password")}
          password={true}
        />
        <View style={styles.newsletter}>
          <RCheckbox />
          <RText text={"  " + I18n.t("newsletterSignup")} accent={true} />
        </View>
        <RButton text={I18n.t("signup")} style={styles.button} />
        <StatusBar style="inverted" />
      </View>
    </>
  );
};

export default SignupScreen;
