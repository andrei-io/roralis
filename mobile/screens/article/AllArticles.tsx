import { ScreenParamsList } from '@/router/router';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { FC } from 'react';
import { StyleSheet } from 'react-native';
import WebView from 'react-native-webview';

type IAllArticlesProps = NativeStackScreenProps<ScreenParamsList, 'AllArticles'>;

const styles = StyleSheet.create({
  container: {
    width: '100%',
    height: '100%',
  },
});

const AllArticlesScreen: FC<IAllArticlesProps> = ({ navigation }) => {
  // TODO: trebuie organizate articolele intr-un blogging engine sau sa vad daca le pot extrage din wordpress
  // deocamdata ramane cu un webview direct catre site
  return (
    <WebView
      style={styles.container}
      source={{ uri: 'https://countryroad.giveitback.ro/category/informatii-utile/' }}
    />
  );
};

export default AllArticlesScreen;
