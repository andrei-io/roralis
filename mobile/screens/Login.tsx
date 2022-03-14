import { ParamsList } from '@/router/router';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import React from 'react';
import { Button, Text, View } from 'react-native';

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
