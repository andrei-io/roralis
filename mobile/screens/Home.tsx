import { isLoggedIn } from '@/cache/auth';
import Colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import { FC, useEffect } from 'react';
import { Image, StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { ScreenParamsList } from '../router/router';

type IHomeProps = NativeStackScreenProps<ScreenParamsList, 'Home'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: Colors.dark.background,
  },
  image: {
    width: '50%',
    height: '30%',
    resizeMode: 'contain',
  },
});

const HomeScreen: FC<IHomeProps> = ({ navigation }) => {
  useEffect(() => {
    isLoggedIn().then(v => {
      if (v) navigation.navigate('AllPosts');
      else navigation.navigate('Landing');
    });
  }, []);
  return (
    <SafeAreaView style={styles.container}>
      <Image source={require('@/assets/logo.png')} style={styles.image} />
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default HomeScreen;
