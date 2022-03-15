import { RPostNormal } from '@/components/core/Post';
import Colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { ScrollView, StyleSheet } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { ParamsList } from '../router/router';

type IDevProps = NativeStackScreenProps<ParamsList, 'Home'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    backgroundColor: Colors.dark.background,
  },
  scollView: {
    flex: 1,
    width: '100%',
  },
});

const DevScreen: React.FC<IDevProps> = ({ navigation }) => {
  return (
    <SafeAreaView style={styles.container}>
      <ScrollView style={styles.scollView}>
        <RPostNormal
          title="Header"
          description="Lorem ipsum dolor sit amet consectetur adipisicing elit. Quibusdam, doloremque."
        />
      </ScrollView>
      {/* <RTextInput placeholder="Nume" /> */}
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default DevScreen;
