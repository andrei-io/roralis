import React from 'react';
import { View, Text, Button, Pressable } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { ParamsList } from '../router/router';

interface IHomeProps {}
type RouterProps = NativeStackScreenProps<ParamsList, 'Home'>;

const HomeScreen: React.FC<IHomeProps & RouterProps> = ({ navigation }) => {
  return (
    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
      <Text>Home Screen</Text>
      <Button
        onPress={() => {
          navigation.push('Login');
        }}
        title="Go to Login"
      />
    </View>
  );
};

export default HomeScreen;
