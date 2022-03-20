import { RText } from '@/components/ui/Text';
import Colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { StyleSheet } from 'react-native';
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
});

const HomeScreen: React.FC<IHomeProps> = ({ navigation }) => {
  return (
    <SafeAreaView style={styles.container}>
      <RText text="Ecran Acasa" accent={true} />
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default HomeScreen;
