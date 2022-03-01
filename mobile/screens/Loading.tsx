import React from 'react';
import { StyleSheet } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { SafeAreaView } from 'react-native-safe-area-context';

import { ParamsList } from '../router/router';
import Colors from '@shared/colors';
import { StatusBar } from 'expo-status-bar';

type IHomeProps = NativeStackScreenProps<ParamsList, 'Loading'>;

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
