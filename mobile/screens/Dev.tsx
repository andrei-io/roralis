import { RCheckbox } from '@components/ui/Checkbox';
import { RText } from '@components/ui/Text';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import Colors from '@shared/colors';
import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { StyleSheet } from 'react-native';
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
});

const DevScreen: React.FC<IDevProps> = ({ navigation }) => {
  return (
    <SafeAreaView style={styles.container}>
      <RCheckbox />
      <RText
        text="Aș dori să primesc buletinul dumneavoastră informativ și alte informații promoționale."
        accent={true}
      />
      <StatusBar style="inverted" />
    </SafeAreaView>
  );
};

export default DevScreen;
