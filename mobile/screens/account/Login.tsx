import { setToken, setUserCache } from '@/cache/auth';
import { RButton } from '@/components/ui/Button';
import { RText } from '@/components/ui/Text';
import { RTextInput } from '@/components/ui/TextInput';
import { SignIn } from '@/restapi/UserAPI';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import { FC, useState } from 'react';
import { Alert, KeyboardAvoidingView, Platform, StyleSheet, View } from 'react-native';
import { SimpleHeader } from '../header/SimpleHeader';

interface ILoginProps {}
type RouterProps = NativeStackScreenProps<ScreenParamsList, 'Login'>;
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
});
const LoginScreen: FC<ILoginProps & RouterProps> = ({ navigation }) => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  async function signIn() {
    if (email == '' || password == '') return;
    try {
      const token = await SignIn(email, password);
      // TODO: make server send back user ID or get user by email
      await setUserCache({ Email: email, Password: password });
      await setToken(token);
      navigation.navigate('AllPosts');
    } catch (e) {
      const error = e as Error;
      Alert.alert(error.message);
    }
  }

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
      enabled={false}
      style={styles.bigContainer}
    >
      <SimpleHeader title={I18n.t('connect')} />
      <View style={styles.container}>
        <RTextInput style={styles.mail} placeholder={I18n.t('email')} onChange={setEmail} />
        <RTextInput
          style={styles.password}
          placeholder={I18n.t('password')}
          password={true}
          onChange={setPassword}
        />
        <RButton text={I18n.t('connect')} style={styles.button} onClick={signIn} />
        <RText text={I18n.t('forgotPassword')} accent={true} />
        <StatusBar style="inverted" />
      </View>
    </KeyboardAvoidingView>
  );
};

export default LoginScreen;
