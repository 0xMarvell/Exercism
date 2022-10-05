class SqueakyClean {
    static String clean(String identifier) {
        StringBuilder sb = new StringBuilder();
        int length = identifier.length();
        for(int i = 0;i<length;i++)
            {
                char ch = identifier.charAt(i);
                if(Character.isSpaceChar(ch) == true)
                    sb.append("_");
                else if(Character.isISOControl(ch) == true)
                    sb.append("CTRL");
                else if(ch == '-'){
                    i++;
                    if(Character.isDigit(identifier.charAt(i)) == true){
                        i++;
                    sb.append(Character.toUpperCase(identifier.charAt(i)));
                    } else {
                    sb.append(Character.toUpperCase(identifier.charAt(i)));
                    }
                }
                else if(Character.isLetter(ch) == false || (ch >= 'α' && ch <= 'ω'))
                    sb.append("");
                else
                    sb.append(ch);
            }

        return sb.toString();
    }
}
