Vim�UnDo� ��26�D/}d��ֶG^�h��K<t1b/�%  J               2   �       X   �   �    V�C    _�       3           2   I        ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          <<<<<<< HEAD5�_�   2   4           3   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          $    /*res, err = taps[3].Get("keyB")5�_�   3   5           4   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          +    if (err != nil) || string(res) != "B" {5�_�   4   6           5   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          H        t.Errorf("TestStoreOracle: Error getting %v: %v\n", "keyB", err)5�_�   5   7           6   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I              }*/5�_�   6   8           7   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          M    /*rootOfA, err:= taps[0].local.findRoot(taps[0].local.node, Hash("keyA"))5�_�   7   9           8   I        ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I           5�_�   8   :           9   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I              for _, tap := range taps{5�_�   9   ;           :   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          &        if tap.local.node == rootOfA {5�_�   :   <           ;   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I                      tap.Kill()5�_�   ;   =           <   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I                      taps5�_�   <   >           =   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          	        }5�_�   =   ?           >   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I              }*/5�_�   >   @           ?   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          0    taps[0].Leave() //tap[0] is the root of keyA5�_�   ?   A           @   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I               fmt.Printf("killing root\n")5�_�   @   B           A   I       ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          /    time.Sleep(2 * REPUBLISH + 2 * time.Second)5�_�   A   C           B   I        ����                                                                                                                                                                                                                                                                                                                                                             V�?�     �   H   I          =======5�_�   B   D           C   M        ����                                                                                                                                                                                                                                                                                                                                                             V�?�   	 �   L   M          0>>>>>>> ff1c489a661365a1d012e50630bc08c217fd2a175�_�   C   E           D   .        ����                                                                                                                                                                                                                                                                                                                                                             V�@     �   -   /  	          }5�_�   D   F           E   -        ����                                                                                                                                                                                                                                                                                                                                                             V�@     �   ,   .            }5�_�   E   G           F   -       ����                                                                                                                                                                                                                                                                                                                                                             V�@     �   ,   2        }5�_�   F   H           G   0        ����                                                                                                                                                                                                                                                                                                                                                             V�@     �   /   2         5�_�   G   I           H   0       ����                                                                                                                                                                                                                                                                                                                                                             V�@+     �   /   0              host := Start(0, "")5�_�   H   J           I   0        ����                                                                                                                                                                                                                                                                                                                                                             V�@,     �   /   1         5�_�   I   K           J   /       ����                                                                                                                                                                                                                                                                                                                                                             V�@K     �   .   0        $func TestRemoveBlops(t *testing.T) {5�_�   J   L           K   0       ����                                                                                                                                                                                                                                                                                                                                                             V�@L     �   /   2            5�_�   K   M           L   1       ����                                                                                                                                                                                                                                                                                                                                                             V�@n     �   0   2        *    a := Start(0, host.local.node.Address)5�_�   L   N           M   1        ����                                                                                                                                                                                                                                                                                                                                                             V�@o     �   1   3      �   1   2      5�_�   M   O           N   2       ����                                                                                                                                                                                                                                                                                                                                                             V�@p     �   2   4      �   2   3      5�_�   N   P           O   3       ����                                                                                                                                                                                                                                                                                                                                                             V�@q     �   3   5      �   3   4      5�_�   O   Q           P   1        ����                                                                                                                                                                                                                                                                                                                                                             V�@r     �   0   1           5�_�   P   R           Q   3   *    ����                                                                                                                                                                                                                                                                                                                                                             V�@w     �   2   6        *    a := Start(0, host.local.node.Address)5�_�   Q   S           R   2       ����                                                                                                                                                                                                                                                                                                                                                             V�@|     �   1   3        *    a := Start(0, host.local.node.Address)5�_�   R   T           S   3       ����                                                                                                                                                                                                                                                                                                                                                             V�@~     �   2   4        *    a := Start(0, host.local.node.Address)5�_�   S   U           T   2       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   1   3        *    b := Start(0, host.local.node.Address)5�_�   T   V           U   3       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   2   4        *    c := Start(0, host.local.node.Address)5�_�   U   W           V   1       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   0   2        *    a := Start(0, host.local.node.Address)5�_�   V   X           W   5       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   3   8             �   4   6            5�_�   W   Y           X   5       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   4   6             aa.Store("aa", []byte("aa"))5�_�   X   Z           Y   6       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   5   7             ab.Store("ab", []byte("ab"))5�_�   Y   [           Z   6       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   5   7        !    ab.Store("ab1", []byte("ab"))5�_�   Z   \           [   5       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   4   6        !    aa.Store("aa1", []byte("aa"))5�_�   [   ]           \   7       ����                                                                                                                                                                                                                                                                                                                                                             V�@�     �   6   ?            5�_�   \   ^           ]   >       ����                                                                                                                                                                                                                                                                                                                                                             V�A     �   =   ?            aa.Remove("ab1")5�_�   ]   _           ^   >       ����                                                                                                                                                                                                                                                                                                                                                             V�A     �   =   C            ab.Remove("ab1")5�_�   ^   `           _   >       ����                                                                                                                                                                                                                                                                                                                                                             V�A=     �   =   A            ab.Remove("ab1")5�_�   _   a           `   @   "    ����                                                                                                                                                                                                                                                                                                                                                             V�AP     �   ?   C        "    ac.Store("ac2", []byte("ac2"))5�_�   `   b           a   A       ����                                                                                                                                                                                                                                                                                                                                                             V�AY     �   A   C  !    �   A   B  !    5�_�   a   c           b   B       ����                                                                                                                                                                                                                                                                                                                                                             V�AY     �   B   D  "    �   B   C  "    5�_�   b   d           c   C       ����                                                                                                                                                                                                                                                                                                                                                             V�AZ     �   C   E  #    �   C   D  #    5�_�   c   e           d   B       ����                                                                                                                                                                                                                                                                                                                                                             V�A]     �   A   C  $      "    ac.Store("ac3", []byte("ac3"))5�_�   d   f           e   C       ����                                                                                                                                                                                                                                                                                                                                                             V�A]     �   B   D  $      "    ac.Store("ac3", []byte("ac3"))5�_�   e   g           f   D       ����                                                                                                                                                                                                                                                                                                                                                             V�A_     �   C   E  $      "    ac.Store("ac3", []byte("ac3"))5�_�   f   h           g   B       ����                                                                                                                                                                                                                                                                                                                                                             V�Ac     �   A   C  $      "    ac.Store("ac4", []byte("ac3"))5�_�   g   i           h   C       ����                                                                                                                                                                                                                                                                                                                                                             V�Ad     �   B   D  $      "    ac.Store("ac5", []byte("ac3"))5�_�   h   j           i   D       ����                                                                                                                                                                                                                                                                                                                                                             V�Ae     �   C   E  $      "    ac.Store("ac6", []byte("ac3"))5�_�   i   k           j   0       ����                                                                                                                                                                                                                                                                                                                                                             V�Ah     �   /   2  $          host := Start(0, "")5�_�   j   l           k   5        ����                                                                                                                                                                                                                                                                                                                                                             V�An     �   4   7  %       5�_�   k   m           l   >        ����                                                                                                                                                                                                                                                                                                                                                             V�As     �   =   @  &       5�_�   l   n           m   B        ����                                                                                                                                                                                                                                                                                                                                                             V�A{     �   A   D  '       5�_�   m   o           n   H   "    ����                                                                                                                                                                                                                                                                                                                                                             V�A     �   G   M  (      "    ac.Store("ac6", []byte("ac6"))5�_�   n   p           o   Q   +    ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   P   T  ,      +    bc := Start(0, host.local.node.Address)5�_�   o   q           p   N        ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   M   O  .       5�_�   p   r           q   S        ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   R   V  .       5�_�   q   s           r   R        ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   Q   T  0       5�_�   r   t           s   U       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   T   X  1          aa.Leave()5�_�   s   u           t   W       ����                                                                                                                                                                                                                                                                                                                                                             V�A�   
 �   V   ]  3          ac.Leave()5�_�   t   v           u   K       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   J   L  8      %    aa.Store("aa1", []bytes("aa1.2"))5�_�   u   w           v   L       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   K   M  8      %    aa.Store("ab1", []bytes("ab1.2"))5�_�   v   x           w   0       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   /   1  8          host := Start(0, "")5�_�   w   y           x   2       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   1   3  8      +    aa := Start(0, host.local.node.Address)5�_�   x   z           y   3       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   2   4  8      +    ab := Start(0, host.local.node.Address)5�_�   y   {           z   4       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   3   5  8      +    ac := Start(0, host.local.node.Address)5�_�   z   |           {   O       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   N   P  8      +    ba := Start(0, host.local.node.Address)5�_�   {   }           |   P       ����                                                                                                                                                                                                                                                                                                                                                             V�A�     �   O   Q  8      +    bb := Start(0, host.local.node.Address)5�_�   |   ~           }   Q       ����                                                                                                                                                                                                                                                                                                                                                             V�A�    �   P   R  8      +    bc := Start(0, host.local.node.Address)5�_�   }              ~   Q   .    ����                                                                                                                                                                                                                                                                                                                                                             V�B     �   P   V  8      .    bc, _ := Start(0, host.local.node.Address)5�_�   ~   �              R        ����                                                                                                                                                                                                                                                                                                                                                             V�B*     �   Q   T  <       5�_�      �           �   V        ����                                                                                                                                                                                                                                                                                                                                                             V�B.    �   U   \  =       5�_�   �   �           �   T   !    ����                                                                                                                                                                                                                                                                                                                                                             V�B�     �   S   U  B      !    ba.Store("ba1", []byte("ba1")5�_�   �   �           �   U   !    ����                                                                                                                                                                                                                                                                                                                                                             V�B�    �   T   V  B      !    ba.Store("ba2", []byte("ba2")5�_�   �   �           �   b        ����                                                                                                                                                                                                                                                                                                                                                             V�B�    �   a   d  B       5�_�   �   �           �   c   4    ����                                                                                                                                                                                                                                                                                                                                                             V�B�    �   b   f  C      4    // A transfer should have happened by this point5�_�   �   �           �   e       ����                                                                                                                                                                                                                                                                                                                                                             V�B�    �   d   f  E          bc.Get("aa1")5�_�   �   �           �   e       ����                                                                                                                                                                                                                                                                                                                                                             V�B�     �   d   i  E          _, err := bc.Get("aa1")5�_�   �   �           �   g        ����                                                                                                                                                                                                                                                                                                                                                             V�B�    �   f   h  H       5�_�   �   �           �   d        ����                                                                                                                                                                                                                                                                                                                                                             V�C     �   c   f  H       5�_�   �               �   d        ����                                                                                                                                                                                                                                                                                                                                                             V�C    �   c   f  I       5��