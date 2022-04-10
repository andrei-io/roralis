import { iconSize, RUserPhoto } from '@/components/auth/UserPhoto';
import { RPostNormal } from '@/components/social/Post';
import { RButton } from '@/components/ui/Button';
import { RNavigationBar } from '@/components/ui/NavigationBar';
import { RText } from '@/components/ui/Text';
import { GetAllPosts, Post } from '@/restapi/PostAPI';
import { GetOneUser, User } from '@/restapi/UserAPI';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import React, { FC, useEffect, useState } from 'react';
import { Pressable, ScrollView, StyleSheet, View } from 'react-native';
import { ProfileHeader } from './ProfileHeader';

type IResetPasswordProps = NativeStackScreenProps<ScreenParamsList, 'Profile'>;

const styles = StyleSheet.create({
  container: {
    width: '100%',
    height: '77%',
    alignItems: 'center',
    backgroundColor: colors.dark.background,
  },
  separator: {
    width: '100%',
    height: '30%',
    backgroundColor: colors.dark.accent,
  },
  content: {
    top: (-3 / 4) * iconSize,
    width: '100%',
    alignItems: 'center',
  },
  posts: {
    paddingTop: 20,
    width: '100%',
    height: '100%',
  },
});

const ProfileScreen: FC<IResetPasswordProps> = ({ navigation, route }) => {
  // TODO: store user info in local data
  const userID = route.params?.id ?? 1;
  const [user, setUser] = useState<User>();
  const [posts, setPosts] = useState<Post[]>();
  useEffect(() => {
    async function fetchData() {
      try {
        const u = await GetOneUser(userID);
        setUser(u);
        const [ps] = await GetAllPosts();
        setPosts(Array(20).fill(ps));
      } catch (e) {
        console.log(e);
      }
    }
    fetchData();
  }, []);
  return (
    <>
      <ProfileHeader />
      <View style={styles.container}>
        <View style={styles.separator} />
        <RUserPhoto name={user?.Name ?? 'Vrooooooooom'} />
        <View style={styles.content}>
          <RText text={user?.Name ?? 'User'} accent={true} size="large" variant="semiBold" />
          <RButton text={I18n.t('createNewPost')} style={{ marginVertical: 20 }} />
          <ScrollView style={styles.posts}>
            {posts?.map((post, i) => (
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
                      (Date.now() - new Date(post.CreatedAt ?? '').getTime()) /
                        (1000 * 60 * 60 * 24),
                    ).toString() + ' z'
                  }
                />
              </Pressable>
            ))}
          </ScrollView>
        </View>
        <StatusBar style="dark" />
      </View>
      <RNavigationBar navigation={navigation} focused="Profile" />
    </>
  );
};

export default ProfileScreen;
