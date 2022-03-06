import colors from '@shared/colors';
import Checkbox from 'expo-checkbox';
import React, { useState } from 'react';

interface ICheckboxProps {
  initial?: boolean;
  onChange?(value: boolean): void;
}

export const RCheckbox: React.FC<ICheckboxProps> = ({ initial = false, onChange }) => {
  const [checked, setChecked] = useState(initial);

  return (
    <Checkbox
      value={checked}
      style={{
        borderRadius: 5,
      }}
      color={checked ? colors.dark.accent : colors.dark.lightGray}
      onValueChange={() => {
        setChecked(!checked);
        onChange?.(!checked);
      }}
    />
  );
};
