import { getToken } from '@/cache/auth';
import { RPhotoAdder } from '@/components/social/PhotoAdder';
import { RButton } from '@/components/ui/Button';
import { RText } from '@/components/ui/Text';
import { RTextInput } from '@/components/ui/TextInput';
import { CreatePost, Post } from '@/restapi/PostAPI';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import React, { FC, useState } from 'react';
import { Alert, StyleSheet, View } from 'react-native';
import { SimpleHeader } from '../header/SimpleHeader';

interface INewPostProps {}
type RouterProps = NativeStackScreenProps<ScreenParamsList, 'NewPost'>;
const styles = StyleSheet.create({
  container: {
    backgroundColor: colors.dark.background,
    flex: 1,
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'space-evenly',
  },
  button: {
    width: '90%',
  },
  title: {
    width: '90%',
    flexDirection: 'column',
    alignSelf: 'center',
    height: '10%',
  },
  description: {
    width: '90%',
    flexDirection: 'column',
    alignSelf: 'center',
    height: '10%',
  },
  photo: {
    width: '90%',
    flexDirection: 'column',
    alignSelf: 'center',
    height: '10%',
  },
});
const NewPostScreen: FC<INewPostProps & RouterProps> = ({ navigation }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const onSubmit = async () => {
    const aborter = new AbortController();
    if (title.length < 2) {
      Alert.alert(I18n.t('titleNotValid'));
      return;
    }
    const jwt = await getToken();

    let post: Post = {};
    post.Title = title;
    post.Description = description;
    post.CategoryID = 1;
    post.RegionID = 1;
    try {
      await CreatePost(post, jwt, aborter);
      navigation.pop();
    } catch (error) {
      const e = error as Error;
      Alert.alert(e.message);
    }
    return () => aborter.abort();
  };
  return (
    <>
      <SimpleHeader title={I18n.t('newPost')} />
      <View style={styles.container}>
        <View style={styles.title}>
          <RText text={I18n.t('title')} style={{ color: colors.dark.white }} variant="medium" />
          <RTextInput placeholder={I18n.t('insertTitle')} onChange={setTitle} />
        </View>
        <View style={styles.description}>
          <RText
            text={I18n.t('description')}
            style={{ color: colors.dark.white }}
            variant="medium"
          />
          <RTextInput placeholder={I18n.t('insertDescription')} onChange={setDescription} />
        </View>
        <View style={styles.photo}>
          <RText
            text={I18n.t('insertImage')}
            style={{ color: colors.dark.white }}
            variant="medium"
          />
          <RPhotoAdder />
        </View>
        <RButton text={I18n.t('createPost')} style={styles.button} onClick={onSubmit} />
        <StatusBar style="inverted" />
      </View>
    </>
  );
};

export default NewPostScreen;
