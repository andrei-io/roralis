import React from 'react';
import { View, Text, Button, Pressable } from 'react-native';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { ParamsList } from '../router/router';

interface ILoginProps {}
type RouterProps = NativeStackScreenProps<ParamsList, 'Home'>;

const LoginScreen: React.FC<ILoginProps & RouterProps> = ({ navigation }) => {
  return (
    <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
      <Text>Login Screen</Text>
      <Button
        onPress={() => {
          navigation.push('Home');
        }}
        title="Go to Home"
      />
    </View>
  );
};

export default LoginScreen;
