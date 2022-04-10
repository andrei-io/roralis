// Generate a hex code from an arbitrary string
// https://stackoverflow.com/questions/3426404/create-a-hexadecimal-colour-based-on-a-string-with-javascript
export function StringToColor(seed: string): string {
  var hash = 0;
  for (var i = 0; i < seed.length; i++) {
    hash = seed.charCodeAt(i) + ((hash << 5) - hash);
  }
  var colour = '#';
  for (var i = 0; i < 3; i++) {
    var value = (hash >> (i * 8)) & 0xff;
    colour += ('00' + value.toString(16)).substr(-2);
  }
  return colour;
}
