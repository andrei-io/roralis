import { RText } from '@/components/ui/Text';
import { GetOnePost, Post } from '@/restapi/PostAPI';
import { GetOneUser, User } from '@/restapi/UserAPI';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import { FC, useEffect, useState } from 'react';
import { StyleSheet, View } from 'react-native';
import { SimpleHeader } from '../header/SimpleHeader';

type IFullPostProps = NativeStackScreenProps<ScreenParamsList, 'OnePost'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingHorizontal: 20,
    // alignItems: 'center',
    backgroundColor: colors.dark.background,
  },
  author: {
    color: colors.dark.white,
    marginBottom: 10,
  },
  title: {
    color: colors.dark.white,
    marginBottom: 10,
  },
  description: {
    color: colors.dark.white,
  },
});

const FullPostScreen: FC<IFullPostProps> = ({ route }) => {
  const [post, setPost] = useState<Post>();
  const [user, setUser] = useState<User>();
  useEffect(() => {
    async function fetchData() {
      try {
        const p = await GetOnePost(route.params.id);
        setPost(p);
        const u = await GetOneUser(p.ID || -1);
        setUser(u);
      } catch (e) {
        console.log(e);
      }
    }
    fetchData();
  }, []);
  return (
    <>
      <SimpleHeader title={I18n.t('postDetails')} />
      <View style={styles.container}>
        <RText
          text={`${I18n.t('author')}: ${user?.Name}`}
          style={styles.author}
          variant="semiBold"
        />
        <RText
          text={`${I18n.t('title')}: ${post?.Description}`}
          style={styles.title}
          variant="medium"
          size="semiLarge"
        />
        <RText
          text={`${I18n.t('description')}: 
${post?.Description}
				`}
          variant="medium"
          style={styles.description}
        />
      </View>
      <StatusBar style="inverted" />
    </>
  );
};

export default FullPostScreen;
