FIRST = %w[
  ascûd[steel] beor[bear] drâth[tree] dûnost[beard] edaris[blood] erôth[earth] fel[frost] hreth[shadow] isidar[rose] jurgen[dragon] knurl[stone] korda[hammer] mithrim[star] orn[eagle] ragni[river] sesti[holy] thorv[shield]
]

SECOND = %w[
  astim[sight] borith[chief] drâth[tree] dûnost[beard] edaris[blood] erôth[earth] fel[frost] gaml[hand] grimst[house] heim[helm] hiem[head] hreth[shadow] isidar[rose] jurgen[dragon] knurl[stone] korda[hammer] kóstha[foot] mithrim[star] nal[hail] nien[heart] ragni[river] sesti[holy] sweld[tear] thorv[shield] thrond[eye] und[warrior]
]

def mix
  names = []
  FIRST.each do |a|
    SECOND.each do |b|
      names << a.capitalize+b
    end
  end
  names
end
