import {
  Inter_400Regular,
  Inter_500Medium,
  Inter_600SemiBold,
  Inter_700Bold,
  useFonts,
} from '@expo-google-fonts/inter';
import { Router } from './router/router';
import { setup } from './translation/setup';

export default function App() {
  let [fontsLoaded] = useFonts({
    Inter_600SemiBold,
    Inter_500Medium,
    Inter_400Regular,
    Inter_700Bold,
  });
  setup();

  if (!fontsLoaded) {
    return null;
  }
  return <Router />;
}
