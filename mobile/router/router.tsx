import LandingScreen from "@/screens/Landing";
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
};

// The api for tabsNavigator is really idiotic so it's easier to hand-roll a component than deal with it
const Stack = createNativeStackNavigator();
export function Router() {
  const production = Constants.manifest?.extra?.production ?? true;

  return (
    <NavigationContainer>
      <Stack.Navigator
        screenOptions={{
          headerShown: false,
          animation: "none",
        }}
        initialRouteName={production ? "Home" : "Login"}
      >
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen name="Login" component={LoginScreen} />
        <Stack.Screen name="Loading" component={LoadingScreen} />
        <Stack.Screen name="Dev" component={DevScreen} />
        <Stack.Screen name="Landing" component={LandingScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}
