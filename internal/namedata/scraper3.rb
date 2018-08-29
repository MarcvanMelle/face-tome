require 'nokogiri'
require 'open-uri'
require 'pry'

def parse_site(language, gender)
  characters = %w[
    a b c d e f g h i j k l m n o p q r s t u v w x y z
  ]

  first_names = []
  uri = 'https://www.babynameguide.com/categoryarabic.asp?strAlpha=[letter]&strCat=[language]&strGender=[gender]'
  uri = uri.gsub('[language]', language)
  uri = uri.gsub('[gender]', gender)
  characters.each do |letter|
    new_uri = uri.gsub('[letter]', letter)
    puts new_uri
    site = Nokogiri::HTML(open(new_uri))
    anchors = site.css('th.NameColumnsName')

    anchors.each do |anchor|
      first_names << anchor.text
    end
  end
  first_names
end

