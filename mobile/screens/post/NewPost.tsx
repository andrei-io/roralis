import { RPhotoAdder } from '@/components/social/PhotoAdder';
import { RButton } from '@/components/ui/Button';
import { RText } from '@/components/ui/Text';
import { RTextInput } from '@/components/ui/TextInput';
import { ScreenParamsList } from '@/router/router';
import colors from '@/shared/colors';
import { NativeStackScreenProps } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import I18n from 'i18n-js';
import React, { FC } from 'react';
import { StyleSheet, View } from 'react-native';
import { SimpleHeader } from '../header/SimpleHeader';

interface INewPostProps {}
type RouterProps = NativeStackScreenProps<ScreenParamsList, 'NewPost'>;
const styles = StyleSheet.create({
  container: {
    backgroundColor: colors.dark.background,
    flex: 1,
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
  },
  button: {
    width: '90%',
    marginVertical: '20%',
  },
  title: {
    width: '90%',
    flexDirection: 'column',
    alignSelf: 'center',
    marginBottom: '20%',
    height: '10%',
  },
  description: {
    width: '90%',
    flexDirection: 'column',
    alignSelf: 'center',
    marginBottom: '20%',
    height: '10%',
  },
  photo: {
    width: '90%',
    flexDirection: 'column',
    alignSelf: 'center',
    marginBottom: '20%',
    height: '10%',
  },
});
const NewPostScreen: FC<INewPostProps & RouterProps> = ({ navigation }) => {
  return (
    <>
      <SimpleHeader title={I18n.t('newPost')} />
      <View style={styles.container}>
        <View style={styles.title}>
          <RText text={I18n.t('title')} style={{ color: colors.dark.white }} variant="medium" />
          <RTextInput placeholder={I18n.t('insertTitle')} />
        </View>
        <View style={styles.description}>
          <RText
            text={I18n.t('description')}
            style={{ color: colors.dark.white }}
            variant="medium"
          />
          <RTextInput placeholder={I18n.t('insertDescription')} password={true} />
        </View>
        <View style={styles.photo}>
          <RText
            text={I18n.t('insertImage')}
            style={{ color: colors.dark.white }}
            variant="medium"
          />
          <RPhotoAdder />
        </View>
        <RButton text={I18n.t('createPost')} style={styles.button} />
        <StatusBar style="inverted" />
      </View>
    </>
  );
};

export default NewPostScreen;
