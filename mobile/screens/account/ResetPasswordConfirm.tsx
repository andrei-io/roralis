import { RVerificationCode } from '@/components/auth/VerificationCode';
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

type IResetPasswordCodeProps = NativeStackScreenProps<ScreenParamsList, 'ResetPasswordCode'>;

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
    marginBottom: 50,
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
});

const ResetPasswordCodeScreen: FC<IResetPasswordCodeProps> = ({ navigation }) => {
  return (
    <>
      <SimpleHeader title={I18n.t('resetPassword')} />
      <View style={styles.container}>
        <RText text={I18n.t('youWillReceiveACode')} variant="medium" style={styles.text} />
        <RTextInput placeholder="Email" style={styles.emailInput} />
        <RButton text={I18n.t('generateCode')} style={styles.button} />
        <RVerificationCode onChange={() => {}} style={styles.verification} />
        <RButton text={I18n.t('verify')} style={styles.button} />

        <StatusBar style="inverted" />
      </View>
    </>
  );
};

export default ResetPasswordCodeScreen;
