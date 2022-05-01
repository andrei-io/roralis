import { RButton } from '@/components/ui/Button';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import { FC } from 'react';
import { Image, StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { ScreenParamsList } from '../router/router';
type ILandingProps = NativeStackScreenProps<ScreenParamsList, 'Landing'>;

const styles = StyleSheet.create({
  container: {
    backgroundColor: colors.dark.background,
    flex: 1,
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'space-evenly',
  },
  button: {
    width: '90%',
  },
  image: {
    width: '50%',
    height: '30%',
    resizeMode: 'contain',
  },
});

const LandingScreen: FC<ILandingProps> = ({ navigation }) => {
  return (
    <SafeAreaView style={styles.container}>
      <Image source={require('@/assets/logo.png')} style={styles.image} />
      <RButton
        text={I18n.t('createAccount')}
        style={styles.button}
        onClick={() => navigation.navigate('Signup')}
      />
      <RButton
        text={I18n.t('connect')}
        style={styles.button}
        onClick={() => navigation.navigate('Login')}
      />
      <RButton
        text={I18n.t('noAccount')}
        style={styles.button}
        onClick={() => navigation.navigate('AllPosts')}
      />
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default LandingScreen;
