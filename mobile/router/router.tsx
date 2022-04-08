import ResetPasswordScreen from "@/screens/account/ResetPassword";
import ResetPasswordCodeScreen from "@/screens/account/ResetPasswordConfirm";
import LandingScreen from "@/screens/Landing";
import AllPostsScreen from "@/screens/post/AllPosts";
import FullPostScreen from "@/screens/post/FullPost";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import Constants from "expo-constants";
import DevScreen from "../screens/Dev";
import HomeScreen from "../screens/Home";
import LoadingScreen from "../screens/Loading";
import LoginScreen from "../screens/Login";

export type ScreenParamsList = {
  Home: undefined;
  Login: undefined;
  Loading: undefined;
  Dev: undefined;
  Landing: undefined;
  ResetPasswordCode: undefined;
  ResetPassword: undefined;
  AllPosts: undefined;
  OnePost: { id: number };
};

// The api for tabsNavigator is really idiotic so it's easier to hand-roll a component than deal with it
const Stack = createNativeStackNavigator();
export function Router() {
  const initialScreenName =
    Constants.manifest?.extra?.initialScreenName ?? "Home";

  return (
    <NavigationContainer>
      <Stack.Navigator
        screenOptions={{
          headerShown: false,
          animation: "none",
        }}
        initialRouteName={initialScreenName}
      >
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen name="Login" component={LoginScreen} />
        <Stack.Screen name="Loading" component={LoadingScreen} />
        <Stack.Screen name="Dev" component={DevScreen} />
        <Stack.Screen name="Landing" component={LandingScreen} />
        <Stack.Screen
          name="ResetPasswordCode"
          component={ResetPasswordCodeScreen}
        />
        <Stack.Screen name="ResetPassword" component={ResetPasswordScreen} />
        <Stack.Screen name="AllPosts" component={AllPostsScreen} />
        <Stack.Screen name="OnePost" component={FullPostScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}
