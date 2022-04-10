import { RPostNormal } from '@/components/social/Post';
import { RTextInput } from '@/components/ui/TextInput';
import { GetAllPosts, Post } from '@/restapi/PostAPI';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import { FC, useEffect, useState } from 'react';
import { Pressable, ScrollView, StyleSheet, View } from 'react-native';
import { SimpleHeader } from '../header/SimpleHeader';

type IAllPostProps = NativeStackScreenProps<ScreenParamsList, 'AllPosts'>;

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    backgroundColor: colors.dark.background,
  },
  searchBar: {
    width: '95%',
    borderRadius: 100,
  },
  postsContainer: {
    flex: 1,
    paddingTop: 20,
    width: '100%',
  },
});

const AllPostsScreen: FC<IAllPostProps> = ({ navigation }) => {
  const [posts, setPosts] = useState<Post[]>([]);
  useEffect(() => {
    GetAllPosts()
      .then(setPosts)
      .catch(e => {
        console.log(e);
      });
  }, []);
  return (
    <>
      <SimpleHeader title={I18n.t('posts')} />
      <View style={styles.container}>
        <RTextInput style={styles.searchBar} placeholder={I18n.t('search')} />
        <ScrollView style={styles.postsContainer}>
          {posts.map((post, i) => (
            <Pressable
              android_disableSound={true}
              onPress={() => {
                navigation.push('OnePost', { id: post.ID ?? 1 });
              }}
              key={i}
            >
              <RPostNormal
                title={post.Title ?? ''}
                description={post.Description}
                time={
                  Math.floor(
                    // TODO: calculeaza corect astea ca nu merge cu time-zone
                    (Date.now() - new Date(post.CreatedAt ?? '').getTime()) / (1000 * 60 * 60 * 24),
                  ).toString() + ' z'
                }
              />
            </Pressable>
          ))}
        </ScrollView>
      </View>
      <StatusBar style="inverted" />
    </>
  );
};

export default AllPostsScreen;
