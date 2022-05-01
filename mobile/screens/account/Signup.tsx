import { setToken, setUserCache } from '@/cache/auth';
import { RButton } from '@/components/ui/Button';
import { RCheckbox } from '@/components/ui/Checkbox';
import { RText } from '@/components/ui/Text';
import { RTextInput } from '@/components/ui/TextInput';
import { SignUp } from '@/restapi/UserAPI';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { checkEmail } from '@/shared/emailChecker';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import { FC, useState } from 'react';
import { Alert, KeyboardAvoidingView, Platform, StyleSheet, View } from 'react-native';
import { SimpleHeader } from '../header/SimpleHeader';

interface ISignuprops {}
type RouterProps = NativeStackScreenProps<ScreenParamsList, 'Signup'>;
const styles = StyleSheet.create({
  bigContainer: {
    width: '100%',
    height: '100%',
  },
  container: {
    backgroundColor: colors.dark.background,
    flex: 1,
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
  },
  button: {
    width: '90%',
    marginVertical: '20%',
  },
  mail: {
    width: '90%',
    marginBottom: '20%',
    height: '6%',
  },
  password: {
    width: '90%',
    marginBottom: '20%',
    height: '6%',
  },
  newsletter: {
    flexDirection: 'row',
  },
});
const SignupScreen: FC<ISignuprops & RouterProps> = ({ navigation }) => {
  const [email, setEmail] = useState<string>('');
  const [name, setName] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  async function signUp() {
    const aborter = new AbortController();
    const errorMessages: string[] = [];
    if (!checkEmail(email)) errorMessages.push(I18n.t('invalidEmail'));
    if (password.length < 6) errorMessages.push(I18n.t('passwordTooShort'));
    if (name.length < 3) errorMessages.push(I18n.t('nameTooShort'));

    if (errorMessages.length != 0) {
      Alert.alert(
        I18n.t('problemCreatingAccount'),
        errorMessages.reduce((final, e) => final + ` - ${e}\n`, ''),
      );
      return;
    }
    try {
      let { Token, User } = await SignUp(name, email, password, aborter);
      if (!User) User = {};
      await setToken(Token ?? '');

      User.Password = password;
      await setUserCache(User ?? {});
      navigation.navigate('AllPosts');
    } catch (e) {
      const error = e as Error;
      Alert.alert(error.message);
    }
    return () => aborter.abort();
  }
  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      enabled={false}
      style={styles.bigContainer}
    >
      <SimpleHeader title={I18n.t('signup')} />
      <View style={styles.container}>
        <RTextInput style={styles.mail} placeholder={I18n.t('name')} onChange={setName} />
        <RTextInput style={styles.mail} placeholder={I18n.t('email')} onChange={setEmail} />
        <RTextInput
          style={styles.password}
          placeholder={I18n.t('password')}
          password={true}
          onChange={setPassword}
        />
        <View style={styles.newsletter}>
          <RCheckbox />
          <RText text={'  ' + I18n.t('newsletterSignup')} accent={true} />
        </View>
        <RButton text={I18n.t('signup')} style={styles.button} onClick={signUp} />
        <StatusBar style="inverted" />
      </View>
    </KeyboardAvoidingView>
  );
};

export default SignupScreen;
