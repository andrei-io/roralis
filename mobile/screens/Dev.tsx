import { RPhotoAdder } from '@/components/social/PhotoAdder';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { SafeAreaView, StyleSheet } from 'react-native';

type IDevProps = NativeStackScreenProps<ScreenParamsList, 'Dev'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: colors.dark.background,
  },
  scollView: {
    flex: 1,
    width: '100%',
  },
  modal: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: 'transparent',
  },
});

const DevScreen: React.FC<IDevProps> = ({ navigation }) => {
  return (
    <SafeAreaView style={styles.container}>
      <RPhotoAdder />
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default DevScreen;
