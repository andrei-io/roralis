import { RButton } from '@/components/ui/Button';
import { RText } from '@/components/ui/Text';
import { RTextInput } from '@/components/ui/TextInput';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import { FC } from 'react';
import { StyleSheet, View } from 'react-native';
import { SimpleHeader } from '../header/SimpleHeader';

type IResetPasswordProps = NativeStackScreenProps<ScreenParamsList, 'ResetPassword'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'space-evenly',

    backgroundColor: colors.dark.background,
  },
  text: {
    width: '80%',
    color: colors.dark.white,
    marginVertical: 20,
  },
  emailInput: {
    width: '90%',
  },
  button: {
    width: '90%',
  },
  verification: {
    marginVertical: 80,
  },
  separator: {
    width: '100%',
    height: '20%',
  },
});

const ResetPasswordScreen: FC<IResetPasswordProps> = ({ navigation }) => {
  return (
    <>
      <SimpleHeader title={I18n.t('resetPassword')} />
      <View style={styles.container}>
        <RText
          text={I18n.t('insertNewPassword')}
          variant="medium"
          style={styles.text}
          size="large"
        />
        <RTextInput placeholder={I18n.t('insertNewPassword')} style={styles.emailInput} />
        <RTextInput placeholder={I18n.t('confirmNewPassword')} style={styles.emailInput} />
        <RButton text={I18n.t('resetPassword')} style={styles.button} />
        <View style={styles.separator} />

        <StatusBar style="inverted" />
      </View>
    </>
  );
};

export default ResetPasswordScreen;
