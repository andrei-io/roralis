import Colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { ScreenParamsList } from '../router/router';

type IHomeProps = NativeStackScreenProps<ScreenParamsList, 'Loading'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.dark.background,
  },
});

const HomeScreen: React.FC<IHomeProps> = ({ navigation }) => {
  return (
    <SafeAreaView style={styles.container}>
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default HomeScreen;
