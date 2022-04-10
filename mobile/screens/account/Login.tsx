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

interface ILoginProps {}
type RouterProps = NativeStackScreenProps<ScreenParamsList, 'Login'>;
const styles = StyleSheet.create({
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
  return (
    <>
      <SimpleHeader title={I18n.t('connect')} />
      <View style={styles.container}>
        <RTextInput style={styles.mail} placeholder={I18n.t('email')} />
        <RTextInput style={styles.password} placeholder={I18n.t('password')} password={true} />
        <RButton text={I18n.t('connect')} style={styles.button} />
        <RText text={I18n.t('forgotPassword')} accent={true} />
        <StatusBar style="inverted" />
      </View>
    </>
  );
};

export default LoginScreen;
