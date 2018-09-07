FIRST_CONJUNCTION = %w[
  Cel Cele Ciri Cirith Cle El Elb Elbe Ethel Ethil Fal Fan Fin Gal Galad Gil Gilad Ithel Ithil Lara Li Lil Lim Lirra Lis Lith Lo Lolien Lon Lori Loth Lui Lune Mi Mil Mith Na Nali Nar Nil Niri Oi Sel Sele Sil Sild Silv Ta Tal Tar Tas Ui
]

SECOND_CONJUNCTION = %w[
  a al ali an ani ati bala be e el ele elo en ha hal i iba il ili ilui in ir iri isi isil isilvi isui ita ith ithl ithlm ithril itil itir itiri la lah le li lil lis lo lui na ne ni no o ol on ri ril sil silv ta thril tir tiri u ui uil uin
]

THIRD_CONJUNCTION = %w[
  ahad alad ar arien bereth berond beth del dhol dol dur dura duril galad had il ilan ilarian ildur ilia ilien ilirien illirien illui illuim ilom ilond ilorien ilune ir iriath iriel irien iril irion irlond isara isil issilira ithil itholien ithril ithui ithuil itir itirion lan lia lien lirien lom lond lorien lui luim lune olien rian riel rien ril sa sil thui thuil tir tirion ui uil uilond uin
]

def merge_elf_names
  full_names = []
  FIRST_CONJUNCTION.each do |a|
    SECOND_CONJUNCTION.each do |b|
      if a.downcase == b.downcase
        a_b = a
      elsif a[-2..-1] == b[0..1]
        a_b = a[0..-3] + b
      elsif a[-1] == b[0]
        a_b = a[0..-2] + b
      else
        a_b = a+b
      end
      THIRD_CONJUNCTION.each do |c|
        if b.downcase == c.downcase
          next
        elsif a_b[-2..-1] == c[0..1]
          a_b_c = a_b[0..-3] + c
        elsif a_b[-1] == c[0]
          a_b_c = a_b[0..-2] + c
        else
          a_b_c = a_b + c
        end
        full_names << a_b_c if !full_names.include?(a_b_c)
      end
    end
  end

  FIRST_CONJUNCTION.each do |a|
    THIRD_CONJUNCTION.each do |c|
      if a.downcase == c.downcase
        next
      elsif a[-2..-1] == c[0..1]
        a_c = a[0..-3] + c
      else
        a_c = a + c
      end
      full_names << a_c if !full_names.include?(a_c)
    end
  end

  full_names
end
